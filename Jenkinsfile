node {

    stage('Initialize'){
        def dockerHome = tool 'docker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
        sh "echo $PWD"

    }


    stage('Build'){
        sh "docker-compose up"
    }
    }