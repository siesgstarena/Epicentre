module github.com/siesgstarena/epicentre

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/joho/godotenv v1.3.0
    github.com/siesgstarena/epicentre/src/config v0.0.0
    github.com/siesgstarena/epicentre/src/web v0.0.0
)

replace (
    github.com/siesgstarena/epicentre/src/config v0.0.0 => ./
    github.com/siesgstarena/epicentre/src/web v0.0.0 => ./
)