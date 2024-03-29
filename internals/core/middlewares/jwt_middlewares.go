package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Usuario struct {
	gorm.Model
	Usuario      string `json:"usuario"`
	Password     string `json:"password"`
	PasswordHash string `gorm:"-"`
	Autorizado   int    `json:"autorizado"`
	Roles        []UsuariosRoles
	// Area         *domain.EncargadoRoles `gorm:"-"`
	// Pertenece    *domain.EncargadoRoles `gorm:"-"`
}

type UsuariosRoles struct {
	gorm.Model
	UsuarioId int `json:"usuarioId"`
	Usuario   Usuario
	RolId     int `json:"rolId"`
	Rol       CatRoles
	ModuloId  int `json:"moduloId"`
}

type CatRoles struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Modulo      int    `json:"modulo"`
}

type CatModulos struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Activo      int    `json:"activo"`
}

type CurrentUser struct {
	ID           float64
	Usuario      string `json:"usuario"`
	Password     string `json:"password"`
	PasswordHash string
	Valido       bool
}

type UserStruct struct {
	ID      int    `json:"id"`
	Usuario string `json:"usuario"`
	// Tipo      *domain.EncargadoRoles `json:"tipo"`
	// Pertenece *domain.EncargadoRoles `json:"pertenece"`
}

var db *gorm.DB

// Init inicializa el middleware y asigna la conexión de la base de datos
func Init(database *gorm.DB) {
	db = database
}

func User(usuarioId int) (*UserStruct, error) {
	var login Usuario
	var response UserStruct
	if res := db.Where("id = ?", usuarioId).Find(&login); res.Error != nil {
		return nil, res.Error
	}
	response.Usuario = login.Usuario
	response.ID = int(login.ID)
	return &response, nil
}

func Usuarios(c *fiber.Ctx) error {
	var login []*Usuario
	var response []*UserStruct
	if res := db.Raw("SELECT DISTINCT usuarios.* FROM usuarios JOIN usuarios_roles ON usuarios_roles.usuario_id = usuarios.id WHERE usuarios_roles.rol_id in (222,223) AND usuarios_roles.deleted_at IS NULL").Find(&login); res.Error != nil {
		return res.Error
	}
	for _, array := range login {
		response = append(response, &UserStruct{ID: int(array.ID), Usuario: array.Usuario})
	}
	return c.Status(200).JSON(&fiber.Map{
		"data": response,
	})
}

// func EncargadoRoles(usuario_id int) (*domain.EncargadoRoles, error) {
// 	var encargadoRoles *domain.EncargadoRoles
// 	if res := db.Raw("SELECT ur.rol_id as rol, COALESCE(CASE WHEN ca.tipoarea_id = 1 AND ca.padre_id != 208 AND ca.id != 280 THEN ca.id END, CASE WHEN ca2.tipoarea_id = 1 AND ca2.padre_id != 208 AND ca2.id != 280 THEN ca2.id END, CASE WHEN ca3.tipoarea_id = 1 AND ca3.padre_id != 208 AND ca3.id != 280 THEN ca3.id END, CASE WHEN ca4.tipoarea_id = 1 AND ca4.padre_id != 208 AND ca4.id != 280 THEN ca4.id END, CASE WHEN ca5.tipoarea_id = 1 AND ca5.padre_id != 208 AND ca5.id != 280 THEN ca5.id END) AS area_id, COALESCE(CASE WHEN ca.tipoarea_id = 1 AND ca.padre_id != 208 AND ca.id != 280 THEN ca.nombre END, CASE WHEN ca2.tipoarea_id = 1 AND ca2.padre_id != 208 AND ca2.id != 280 THEN ca2.nombre END, CASE WHEN ca3.tipoarea_id = 1 AND ca3.padre_id != 208 AND ca3.id != 280 THEN ca3.nombre END, CASE WHEN ca4.tipoarea_id = 1 AND ca4.padre_id != 208 AND ca4.id != 280 THEN ca4.nombre END, CASE WHEN ca5.tipoarea_id = 1 AND ca5.padre_id != 208 AND ca5.id != 280 THEN ca5.nombre END) AS area FROM personas_usuarios pu INNER JOIN empleados e ON e.persona_id = pu.persona_id INNER JOIN rehu_contratosempleados rc ON rc.empleado_id = e.id INNER JOIN rehu_areasempleados ra ON ra.contratoempleado_id = rc.id INNER JOIN usuarios_roles ur ON ur.usuario_id = ? LEFT JOIN cat_areas ca ON ca.id = ra.area_id LEFT JOIN cat_areas ca2 ON ca2.id = ca.padre_id LEFT JOIN cat_areas ca3 ON ca3.id = ca2.padre_id LEFT JOIN cat_areas ca4 ON ca4.id = ca3.padre_id LEFT JOIN cat_areas ca5 ON ca5.id = ca4.padre_id WHERE pu.usuario_id = ? AND ur.rol_id = 224 AND rc.fecha_termino IS NULL AND ra.fecha_termino IS NULL", usuario_id, usuario_id).Scan(&encargadoRoles); res.Error != nil {
// 		return nil, res.Error
// 	}
// 	return encargadoRoles, nil
// }

func GenerarLogAlumno(c *fiber.Ctx) error {
	// jwt := new(Usuario)
	jwt := new(domains.Alumno)
	if err := c.BodyParser(jwt); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Invalid body parser",
		})
	}
	if jwt.Matricula == "" || jwt.Password == "" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Invalid credentials",
		})
	}

	res, err := GenerateTokenAlumno(jwt)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"usuario": jwt.Matricula,
		"token":   res,
	})
}

func GenerarLogMaestro(c *fiber.Ctx) error {
	// jwt := new(Usuario)
	jwt := new(domains.Personal)
	if err := c.BodyParser(jwt); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Invalid body parser",
		})
	}
	if jwt.User == "" || jwt.Password == "" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Invalid credentials",
		})
	}

	res, perfiles, err := GenerateTokenMaestro(jwt)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"usuario":  jwt.User,
		"token":    res,
		"perfiles": perfiles,
	})
}

func Verificar(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token invalido",
		})
	}
	tokenString := strings.Split(authHeader, " ")[1]
	token, err := ValidateToken(tokenString)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if token == nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token invalido",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"id":      token.ID,
		"usuario": token.Usuario,
		"valido":  token.Valido,
	})
}

func Autorizar(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token inválido",
		})
	}
	tokenString := strings.Split(authHeader, " ")[1]
	token, err := ValidateToken(tokenString)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if token == nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token inválido",
		})
	}

	return c.Next()
	// c.Status(404).JSON(&fiber.Map{
	// 	"message": "no next",
	// })
}

func Refrescar(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token inválido",
		})
	}
	tokenString := strings.Split(authHeader, " ")[1]
	token, err := ValidateToken(tokenString)
	if err != nil {

		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if token == nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token inválido",
		})
	}
	if token.Valido {
		res, err := RefreshToken(tokenString)
		if err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": res,
		})
	} else {
		return c.Status(404).JSON(&fiber.Map{
			"message": "usuario o contraseña incorrectos",
		})
	}
}

func GenerateTokenAlumno(jwtParams *domains.Alumno) (string, error) {
	var login *domains.Alumno
	if res := db.Preload(clause.Associations).Where("matricula = ?", jwtParams.Matricula).First(&login); res.Error != nil {
		return "", res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(jwtParams.Password))
	if err != nil {
		return "", err
	}
	jwtParams.PasswordHash = login.Password
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = login.ID
	claims["userId"] = jwtParams.Matricula
	claims["passwordHash"] = jwtParams.PasswordHash
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte("secret.code"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenMaestro(jwtParams *domains.Personal) (string, []domains.Perfil, error) {
	var login *domains.Personal
	if res := db.Preload(clause.Associations).Preload("Perfiles").Where("user = ?", jwtParams.User).First(&login); res.Error != nil {
		return "", nil, res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(jwtParams.Password))
	if err != nil {
		return "", nil, err
	}
	jwtParams.PasswordHash = login.Password
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = login.ID
	claims["userId"] = jwtParams.User
	claims["passwordHash"] = jwtParams.PasswordHash
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte("secret.code"))
	if err != nil {
		return "", nil, err
	}

	return tokenString, login.Perfiles, nil
}

func ValidateToken(tokenString string) (*CurrentUser, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret.code"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		Id := claims["id"].(float64)
		userId := claims["userId"].(string)
		passwordHash := claims["passwordHash"].(string)

		var login *domains.Alumno
		var login2 *domains.Personal
		var valido bool = false

		res := db.Where("matricula = ? AND password = ? AND id = ?", userId, passwordHash, Id).First(&login)
		if res.RowsAffected == 0 {
			login = nil
			res2 := db.Where("user=? AND password = ? AND id = ?", userId, passwordHash, Id).First(&login2)
			if res2.RowsAffected == 0 {
				login2 = nil
				return nil, errors.New("usuario o contraseña inválidos")
			}
		}

		if login != nil {
			if login.Password != passwordHash {
				return nil, fmt.Errorf("usuario o contraseña inválidos")
			} else {
				valido = true
				return &CurrentUser{Usuario: userId, ID: Id, Valido: valido}, nil
			}
		} else if login2 != nil {
			if login2.Password != passwordHash {
				return nil, fmt.Errorf("usuario o contraseña inválidos")
			} else {
				valido = true
				return &CurrentUser{Usuario: userId, ID: Id, Valido: valido}, nil
			}
		}

	} else {
		return nil, err
	}
	return nil, err
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret.code"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		id := claims["id"]
		userId := claims["userId"].(string)
		passwordHash := claims["passwordHash"].(string)

		var login *domains.Alumno
		var login2 *domains.Personal
		var myFlag bool = false
		var myFlag2 bool = false
		if res := db.Where("matricula = ?", userId).First(&login); res.Error != nil {
			myFlag = true
			if !myFlag {
				return "", res.Error
			}
		} else if myFlag {
			if res2 := db.Where("user=?", userId).First(&login2); res2.Error != nil {
				return "", res2.Error
			}
		}
		if login != nil {
			fmt.Print(login.Matricula)
		} else {
			fmt.Print(login2.User)
		}

		if login.Password != passwordHash {
			myFlag2 = true
			if !myFlag2 {
				return "", fmt.Errorf("token inválido")
			}
		} else if myFlag2 {
			if login2.Password != passwordHash {
				return "", fmt.Errorf("token inválido")
			}

		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		//ID???? Sí le faltaba
		claims["id"] = id
		claims["userId"] = userId
		claims["passwordHash"] = passwordHash
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		tokenString, err := token.SignedString([]byte("secret.code"))
		if err != nil {
			return "", err
		}

		return tokenString, nil
	} else {
		return "", err
	}
}
