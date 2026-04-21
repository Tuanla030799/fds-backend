FROM maven:3.9.10-eclipse-temurin-21 AS build
WORKDIR /app
COPY springboot-backend/pom.xml ./pom.xml
COPY springboot-backend/src ./src
RUN mvn -q -DskipTests package

FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=build /app/target/fds-backend-1.0.0.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/app/app.jar"]
