CREATE TABLE `alumnos` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `matricula` varchar(191) NOT NULL,
  `estado` longtext NOT NULL COMMENT '''(baja_temporal, baja_definitiva, activo)''',
  `password` longtext DEFAULT NULL,
  `fecha_inscripcion` datetime(3) DEFAULT NULL,
  `terminos` bigint(20) NOT NULL,
  `visita_biblioteca` bigint(20) NOT NULL,
  `visita_reglamento` bigint(20) NOT NULL,
  `visita_estadia` bigint(20) NOT NULL,
  `fecha_biblioteca` datetime(3) NOT NULL,
  `fecha_reglamento` datetime(3) NOT NULL,
  `fecha_estadia` datetime(3) NOT NULL,
  `visita_acoso` bigint(20) NOT NULL,
  `fecha_acoso` datetime(3) NOT NULL,
  `visita_igualdad` bigint(20) NOT NULL,
  `fecha_igualdad` datetime(3) NOT NULL,
  `fecha_clases` datetime(3) NOT NULL,
  `visita_clases` bigint(20) NOT NULL,
  `estudiante_id` bigint(20) UNSIGNED DEFAULT NULL,
  `carrera_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `alumno_grupos` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `alumno_id` bigint(20) UNSIGNED NOT NULL,
  `grupo_id` bigint(20) UNSIGNED NOT NULL,
  `estado` longtext DEFAULT NULL COMMENT 'Indica el estado del grupo para el alumno(aprobado/reprobado)',
  `observacion` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `calificaciones` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `grupo_materia_id` bigint(20) UNSIGNED NOT NULL,
  `alumno_id` bigint(20) UNSIGNED NOT NULL,
  `primero` double DEFAULT NULL,
  `segundo` double DEFAULT NULL,
  `tercero` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE `carreras` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nombre` longtext NOT NULL,
  `tipo` longtext NOT NULL,
  `duracion` bigint(20) NOT NULL,
  `descripcion` varchar(191) DEFAULT NULL,
  `estado` tinyint(1) NOT NULL,
  `admision` tinyint(1) NOT NULL,
  `siglas` varchar(191) DEFAULT NULL,
  `nummax1` bigint(20) NOT NULL,
  `nummax2` bigint(20) NOT NULL,
  `nummax3` bigint(20) NOT NULL,
  `utcampus_id` bigint(20) UNSIGNED DEFAULT NULL,
  `personal_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `estudiantes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nombre` longtext NOT NULL COMMENT '''Nombre del aspirante''',
  `apellido_p` longtext NOT NULL COMMENT '''Apellido paterno del aspirante''',
  `apellido_m` varchar(191) DEFAULT NULL COMMENT '''Apellido materno del aspirante''',
  `sexo` longtext NOT NULL COMMENT '''Género del aspirante (M/F)''',
  `discapacidad` varchar(191) DEFAULT NULL COMMENT '''Discapacidad del aspirante''',
  `sangre` varchar(191) DEFAULT NULL COMMENT '''Tipo de sangre del aspirante''',
  `idioma` longtext NOT NULL COMMENT '''Lengua materna del aspirante''',
  `fecha_nacimiento` datetime(3) NOT NULL COMMENT '''Fecha de nacimiento del aspirante''',
  `edad` bigint(20) DEFAULT NULL COMMENT '''Edad del aspirante''',
  `estado_civil` longtext NOT NULL COMMENT '''Estado civil del aspirante (S/C)''',
  `curp` longtext NOT NULL COMMENT '''CURP del aspirante''',
  `nss` longtext NOT NULL COMMENT '''Número de Seguro Social (NSS)''',
  `calle` longtext NOT NULL COMMENT '''Calle de la dirección del aspirante''',
  `numero` varchar(191) DEFAULT NULL COMMENT '''Número del domicilio del aspirante''',
  `colonia` varchar(191) DEFAULT NULL COMMENT '''Colonia donde vive el aspirante''',
  `distrito` longtext NOT NULL COMMENT '''Distrito del aspirante''',
  `ciudad` varchar(191) DEFAULT NULL COMMENT '''Ciudad reconocida del aspirante''',
  `estado` longtext NOT NULL COMMENT '''Entidad federativa del aspirante''',
  `cp` bigint(20) NOT NULL COMMENT '''Código Postal (CP) del aspirante''',
  `localidad` varchar(191) DEFAULT NULL COMMENT '''Localidad del aspirante''',
  `municipio` longtext NOT NULL COMMENT '''Municipio donde reside el aspirante''',
  `email` varchar(191) DEFAULT NULL COMMENT '''Correo electrónico del aspirante''',
  `telefono` varchar(191) DEFAULT NULL COMMENT '''Número telefónico del aspirante''',
  `esc_procedencia` longtext NOT NULL COMMENT '''Nombre de la escuela de procedencia''',
  `subsistema_procedencia` longtext NOT NULL COMMENT '''Subsistema de la escuela de procedencia''',
  `tipo_procedencia` longtext NOT NULL COMMENT '''Tipo de procedencia (Privada/Pública/Cooperación)''',
  `area_procedencia` longtext NOT NULL COMMENT '''Área de bachillerato de procedencia''',
  `fecha_egreso_procedencia` datetime(3) NOT NULL COMMENT '''Fecha de egreso del nivel medio superior''',
  `municipio_procedencia` longtext NOT NULL COMMENT '''Municipio donde se encuentra la escuela de procedencia''',
  `estado_procedencia` longtext NOT NULL COMMENT '''Estado donde se encuentra la escuela de procedencia''',
  `promedio_procedencia` double DEFAULT NULL COMMENT '''Promedio obtenido en el nivel medio superior''',
  `duracion_procedencia` bigint(20) NOT NULL COMMENT '''Duración del plan de estudios del nivel medio superior (1/2)''',
  `secundaria` longtext NOT NULL COMMENT '''Nombre de la secundaria donde cursó''',
  `promedio_secundaria` double DEFAULT NULL COMMENT '''Promedio obtenido en la secundaria''',
  `tutor` varchar(191) DEFAULT NULL COMMENT '''Nombre del tutor''',
  `tutor_ap` longtext NOT NULL COMMENT '''Apellido paterno del tutor''',
  `tutor_am` varchar(191) DEFAULT NULL COMMENT '''Apellido materno del tutor''',
  `direccion_tutor` varchar(191) DEFAULT NULL COMMENT '''Dirección del tutor''',
  `telefono_tutor` varchar(191) DEFAULT NULL COMMENT '''Número telefónico del tutor''',
  `email_tutor` varchar(191) DEFAULT NULL COMMENT '''Correo electrónico del tutor''',
  `observaciones` varchar(191) DEFAULT NULL COMMENT '''Observaciones sobre el aspirante''',
  `obserficha` longtext NOT NULL COMMENT '''Observaciones en ficha''',
  `id_pficha` bigint(20) NOT NULL COMMENT '''ID de la preficha''',
  `fecha_ficha` datetime(3) NOT NULL COMMENT '''Fecha de la ficha''',
  `matricula_ceneval` longtext NOT NULL COMMENT '''Matrícula del examen CENEVAL''',
  `pass_ficha` longtext NOT NULL COMMENT '''Contraseña de la ficha''',
  `car_ficha` bigint(20) NOT NULL COMMENT '''Carrera de la ficha''',
  `car2_ficha` bigint(20) NOT NULL COMMENT '''Segunda opción de carrera de la ficha''',
  `nivel_fi` bigint(20) NOT NULL COMMENT '''Nivel FI''',
  `op_por` longtext NOT NULL COMMENT '''OP POR''',
  `acta` bigint(20) NOT NULL COMMENT '''Acta''',
  `curpp` bigint(20) NOT NULL COMMENT '''CURPP''',
  `identi` bigint(20) NOT NULL COMMENT '''Identificador''',
  `cer` bigint(20) NOT NULL COMMENT '''CER''',
  `curpins` bigint(20) NOT NULL COMMENT '''CURPINS''',
  `certins` bigint(20) NOT NULL COMMENT '''CERTINS''',
  `actains` bigint(20) NOT NULL COMMENT '''ACTAINS''',
  `certmedicoins` bigint(20) NOT NULL COMMENT '''CERTMEDICOINS''',
  `compdomiins` bigint(20) NOT NULL COMMENT '''COMPDOMIINS''',
  `pagoins` bigint(20) NOT NULL COMMENT '''PAGOINS''',
  `des_cermediins` varchar(191) DEFAULT NULL COMMENT '''Des Cermediins''',
  `des_certins` varchar(191) DEFAULT NULL COMMENT '''Des Certins''',
  `des_actains` varchar(191) DEFAULT NULL COMMENT '''Des Actains''',
  `des_comdomiins` varchar(191) DEFAULT NULL COMMENT '''Des Comdomiins''',
  `cartacomp` bigint(20) NOT NULL COMMENT '''Carta de compromiso''',
  `d_pagoins` varchar(191) DEFAULT NULL COMMENT '''D Pagoins''',
  `d_certins` varchar(191) DEFAULT NULL COMMENT '''D Certins''',
  `d_certmedicoins` varchar(191) DEFAULT NULL COMMENT '''D Certmedicoins''',
  `d_actains` varchar(191) DEFAULT NULL COMMENT '''D Actains''',
  `d_compins` varchar(191) DEFAULT NULL COMMENT '''D Compins''',
  `d_curp` varchar(191) DEFAULT NULL COMMENT '''D Curp''',
  `id_preficha` bigint(20) NOT NULL COMMENT '''ID de la preficha'''
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `grupos` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nombre` longtext NOT NULL,
  `cuatrimestre` bigint(20) NOT NULL,
  `carrera_id` bigint(20) UNSIGNED NOT NULL,
  `periodo_id` bigint(20) UNSIGNED NOT NULL,
  `cupo` bigint(20) NOT NULL,
  `personal_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE `grupos_materias` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `grupo_id` bigint(20) UNSIGNED DEFAULT NULL,
  `materia_id` bigint(20) UNSIGNED DEFAULT NULL,
  `personal_id` bigint(20) UNSIGNED DEFAULT NULL,
  `primero` bigint(20) DEFAULT NULL,
  `segundo` bigint(20) DEFAULT NULL,
  `tercero` bigint(20) DEFAULT NULL,
  `acta1` bigint(20) DEFAULT NULL,
  `acta2` bigint(20) DEFAULT NULL,
  `acta3` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `materias` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nombre` longtext NOT NULL,
  `horas` bigint(20) DEFAULT NULL,
  `cuatrimestre` bigint(20) NOT NULL COMMENT '''cuarimestre de la materia''',
  `carrera_id` bigint(20) UNSIGNED NOT NULL COMMENT '''carrera a la que pertenece la materia''',
  `estado` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `perfiles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `perfil` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `periodos` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `inicio_admision` datetime(3) DEFAULT NULL,
  `fin_admision` datetime(3) DEFAULT NULL,
  `fecha_inicio` datetime(3) DEFAULT NULL,
  `fecha_fin` datetime(3) DEFAULT NULL,
  `inicio_reins` datetime(3) DEFAULT NULL,
  `fin_reins` datetime(3) DEFAULT NULL,
  `examenprof` bigint(20) DEFAULT NULL,
  `titulacion` bigint(20) DEFAULT NULL,
  `titulacioning` bigint(20) DEFAULT NULL,
  `primero` tinyint(1) DEFAULT NULL COMMENT 'activar/desactivar primer parcial para calificaciones',
  `segundo` tinyint(1) DEFAULT NULL COMMENT 'activar/desactivar segundo parcial para calificaciones',
  `tercero` tinyint(1) DEFAULT NULL COMMENT 'activar/desactivar tercer parcial para calificaciones',
  `p1` bigint(20) DEFAULT NULL,
  `p2` bigint(20) DEFAULT NULL,
  `p3` bigint(20) DEFAULT NULL,
  `inicio_ev` datetime(3) DEFAULT NULL,
  `fin_ev` datetime(3) DEFAULT NULL,
  `parcial` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE `personal` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `grado` longtext DEFAULT NULL,
  `nombre` longtext DEFAULT NULL,
  `apellido_p` longtext DEFAULT NULL,
  `apellido_m` longtext DEFAULT NULL,
  `user` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `estado` tinyint(1) DEFAULT NULL,
  `adscrito` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE `usuarios_perfiles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `personal_id` bigint(20) UNSIGNED NOT NULL,
  `perfil_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE `ut_campus` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nombre` longtext DEFAULT NULL,
  `siglas` longtext DEFAULT NULL,
  `ubicacion` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

