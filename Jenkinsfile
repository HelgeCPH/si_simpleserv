pipeline {
    agent {
        docker {
            image 'golang:1.11' 
            args '-v "$PWD":/usr/src/minisrv -w /usr/src/minisrv' 
        }
    }
    stages {
        stage('Build') { 
            steps {
                sh 'make all' 
            }
        }
    }
}