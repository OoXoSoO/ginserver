[![Go test](https://github.com/OoXoSoO/ginserver/actions/workflows/pipeline.yml/badge.svg)](https://github.com/OoXoSoO/ginserver/actions/workflows/pipeline.yml)

# GINSERVER

Servidor que permite ejecutar una aplicación utilizando diferentes frameworks para la capa de presentación. Esta está desacoplada completamente del resto del software.
Por otro lado se ha utilizado este repositorio para aprender a utilizar githubactions y [`testcontainers`](https://golang.testcontainers.org/)

## Arquitectura Limpia en la Capa de Presentación

Este proyecto implementa la arquitectura limpia en la capa de presentación con dos enfoques diferentes:

1. Implementación utilizando la runtime de Go.
2. Implementación utilizando el framework Gogin.

### Implementación con la runtime de Go

TODO

### Implementación con Gin

TODO

## GitHub Actions para Pruebas Automáticas

Este proyecto utiliza GitHub Actions para automatizar las pruebas del código. Se han configurado acciones que se ejecutan automáticamente cuando se realiza un push o pull request en el repositorio. Esto garantiza que las pruebas se ejecuten de manera consistente y ayuda a mantener la integridad del código.

### Configuración de GitHub Actions

Este proyecto utiliza GitHub Actions para ejecutar automáticamente pruebas en cada pull request dirigido a la rama `master`. La configuración se encuentra en el archivo [`.github/workflows/pipeline.yml`](.github/workflows/pipeline.yml).

## Requisitos del Sistema

Asegúrate de tener instalado lo siguiente antes de ejecutar el proyecto:

- Go versión 1.20

## Contribuciones

Si deseas contribuir a este proyecto, sigue los pasos a continuación:

1. Haz un fork del repositorio
2. Crea una nueva rama (`git checkout -b feature/nueva-funcionalidad`)
3. Realiza tus cambios y haz commit (`git commit -am 'Agrega nueva funcionalidad'`)
4. Sube los cambios a tu fork (`git push origin feature/nueva-funcionalidad`)
5. Abre un pull request

## Licencia

Este proyecto está bajo la licencia [Nombre de la Licencia]. Consulta el archivo LICENSE.md para obtener más detalles.
