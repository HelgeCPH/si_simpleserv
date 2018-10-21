pipeline {
    agent none
    
    
    stages {
        stage('Build') { 
            agent {
                docker {
                    image 'golang:1.11' 
                    args '-v "$PWD":/usr/src/minisrv -w /usr/src/minisrv' 
                }
            }
            steps {
                sh 'make all' 
            }
        }
        stage('Build') { 
            agent {
                docker {
                    image 'golang:1.11' 
                    args '-v "$PWD":/usr/src/minisrv -w /usr/src/minisrv' 
                }
            }
            steps {
                sh 'make all' 
            }
            post {
                success {
                    archiveArtifacts 'dist/miniserv-*'
                }
            }
        }
    }
}