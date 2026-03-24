# Generador de imagenes QR

Se establece la documentación minima para determinar el contrato del microservicio de generación de códigos QR's

El microservicio debera ser capaz de generar una imagen con un Titulo, un subtitulo, una codigo QR que rediriga a un link pasado por parametro y un mensaje opcional, todo esto en una sola imagen PNG con una relación de aspecto que tienda a ser mas vertical que horizontal (1:1.5 ~ 2).
Dicho servicio, deber permitir personalizar los siguientes aspectos segun variables de entorno:
- Color de fondo
- Color de titulo
- Color de borde
- Color de texto del subtitulo
- Color de texto del mensaje
- Color de texto del titulo
- Color del codigo QR

### Generar imagen

#### Request

* path: `/qr/generate`
* method: `POST`
* body: `IGenerateQrParams`

```typescript
interface IGenerateQrParams{
    link: string
    title: string
    subtitle?: string
    message?: string
}
```

#### Response

* status: `200`
* body: `Buffer`
* contentType: `image/png`

