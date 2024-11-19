# Usa una imagen base de Go para compilar el binario
FROM golang:1.23-alpine AS builder

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias y compila el binario
RUN go mod download
RUN go build -o mc-auth ./cmd/.

# Imagen final para ejecución
FROM alpine:latest

WORKDIR /root/

# Copia el binario desde la etapa de compilación
COPY --from=builder /app/mc-auth .

# Otorga permisos de ejecución al binario
RUN chmod +x /root/mc-auth

# Crear la estructura de directorios y copiar el archivo .env
RUN mkdir -p /root/cmd
COPY --from=builder /app/cmd/.env /root/cmd/.env
COPY --from=builder /app/i18n /root/i18n
COPY --from=builder /app/cmd/cert /root/cmd/cert

# Configura el contenedor para que use el archivo .env
ENV ENV_FILE_PATH=/root/cmd/.env

# Exponer el puerto en el que corre el servicio
EXPOSE 8081

# Comando de inicio
CMD ["sh", "-c", "export $(grep -v '^#' .env | xargs) && ./mc-auth"]
