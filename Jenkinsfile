pipeline {
    environment {
        PW_STAG=credentials('arangoMatchesStag')
        PW_PROD=credentials('arangoMatchesProd')
    }
    agent any

    stages {
        stage ("Prepare stag environment") {
            steps {
                sh "docker stop matches-service-stag &"
                sh "docker stop matches-arangodb-stag &"
                sh "docker rm matches-arangodb-stag &"
                sh "docker rm matches-service-stag &"
                sh "docker kill matches-arangodb-stag &"
                sh "docker kill matches-service-stag &"
            }
        }

        stage ("Build") {
            steps{
                sh "echo PW_STAG=${PW_STAG} >> .env"
                sh "sleep 3s"
                sh "docker-compose -f docker-compose.yml -f docker-compose.stag.yml build"
            }
        }

        stage ("Staging") {
            steps {
                sh "docker-compose -p matches-stag -f docker-compose.stag.yml up -d --force-recreate"
                sh "sleep 60s"
            }
        }

        stage ("Test") {
            steps {
                sh "docker exec matches-service-stag /matches.test --dbhost=arangodb --dbport=8529 --dbpassword=$PW_STAG"
            }
        }

        stage ("Push Staging to Docker Repo") {
            steps {
                sh "docker tag matches-service-stag localhost:5000/matches-service-stag"
                sh "docker push localhost:5000/matches-service-stag"
            }
        }

        stage ("Copy openAPI spec file") {
            steps {
            sh "mkdir ../openAPISpecs"
            sh "cp matches.yml ../openAPISpecs/"
            }
        }

        stage ("Prepare prod environment") {
            steps {
                sh "echo PW_PROD=${PW_PROD} >> .env"
                sh "sleep 3s"
                sh "docker-compose -f docker-compose.yml -f docker-compose.prod.yml build"
                sh "rm -f .env"
            }
        }

        stage ("Production") {
           steps{
                sh "docker stop matches-service-prod &"
                sh "docker stop matches-arangodb-prod &"
                sh "docker rm matches-arangodb-prod &"
                sh "docker rm matches-service-prod &"
                sh "sleep 15s"
                sh "docker-compose -p matches-prod -f docker-compose.prod.yml up --force-recreate"
            }
        }
    }
    post {
        always {
            sh "docker-compose -f docker-compose.stag.yml down -v --rmi 'all'"
            sh "docker-compose -f docker-compose.prod.yml down -v --rmi 'all'"
        }
    }
}