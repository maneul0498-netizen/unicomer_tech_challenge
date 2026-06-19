# unicomer_tech_challenge
Unicomer tech challenge

## 📌 Clonar el repositorio
```bash
git clone https://github.com/maneul0498-netizen/unicomer_tech_challenge.git
```
---
## 📌 Intrucciones de arranque
- Moverse a la raiz del proyecto
```bash
cd unicomer_tech_challenge
```
---
- Ejecutar la construccion de la imagen con docker

```bash
docker build -t unicomer_tech_challenge .
```
---
- Ejecutar la imagen generada con la variable de entorno ADDRESS_HTTP establecida
```bash
docker run -e ADDRESS_HTTP=:8080 -p 8080:8080 unicomer_tech_challenge
```
---
- En el navegador habrir la url http://localhost:8080/api/v1/swagger/index.html para ver la documentacion del endpoint REST

---
- Peticion que acepta json
```bash
curl --location 'http://localhost:8080/api/v1/holidays?fromDate=2026-12-08&toDate=2026-12-31&type=Civil' -H 'Accept: application/json'
```
- Observe que se realiza filtro por fecha y tipo de festivo
---

- Peticion que acepta xml
```bash
curl --location 'http://localhost:8080/api/v1/holidays?fromDate=2026-12-08&toDate=2026-12-31&type=Civil' -H 'Accept: application/xml'
```
- Observe que se realiza filtro por fecha y tipo de festivo
