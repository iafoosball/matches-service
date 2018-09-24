node {

    stage('Initialize'){
        def dockerHome = tool 'docker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
        sh "ls"

    }


    stage('Build'){
        sh "docker-compose up"
    }
    }