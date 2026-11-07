# Sistema de gestión de streaming

Proyecto académico de Programación Orientada a Objetos — UIDE.

**Autor:** Alexis Aguilar
**Estado:** planificación inicial; implementación pendiente.

## Objetivo
Diseñar una aplicación de consola en Go para administrar un catálogo de contenidos audiovisuales, usuarios y favoritos mediante módulos y funciones.

## Alcance
- Administrador: registrar, consultar, editar y desactivar títulos.
- Usuario: consultar y filtrar el catálogo y gestionar favoritos.
- Persistencia local prevista en JSON. No incluye transmisión real de video, pagos ni licencias comerciales.

## Módulos previstos
- `cmd/streamuide`: punto de entrada y menú.
- `internal/catalogo`: validación, búsqueda y mantenimiento de títulos.
- `internal/usuarios`: registro y consulta.
- `internal/favoritos`: listas y comprobación de referencias.
- `internal/almacenamiento`: lectura y escritura de datos.

## Paquetes previstos
Biblioteca estándar: `fmt`, `encoding/json`, `os`, `strings` y `errors`. Cualquier dependencia adicional se documentará antes de incorporarse.

## Próximos pasos
Definir estructuras y contratos de funciones, inicializar el módulo Go con la ruta del repositorio, implementar cada módulo y comprobar su funcionamiento con datos ficticios. El repositorio contiene por ahora únicamente la planificación.

## Referencias
- https://go.dev/doc/tutorial/create-module
- https://go.dev/doc/modules/layout
- https://pkg.go.dev/std
