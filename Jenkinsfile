pipeline {

    agent any

    stages {

        stage('Initialize'){
            steps{
                def dockerHome = tool 'docker'
                env.PATH = "${dockerHome}/bin:${env.PATH}"
            }
        }


        stage('Deploy'){
            echo "$PWD"
            sh "docker-compose up"
        }
    }
}