node {
    def app

    stage('Clone Repository') {
        try {
            checkout scm
        } catch (err) {
            error "Failed to clone repository: ${err}"
        }
    }

    stage('Build Image') {
        try {
            app = docker.build("meshboy/godeploy:${env.BUILD_NUMBER}")
        } catch (err) {
            error "Failed to build Docker image: ${err}"
        }
    }

    stage('Test Image') {
        try {
            app.inside {
                sh 'echo "Running tests inside the container"'
                // Add actual test commands here
            }
        } catch (err) {
            error "Tests failed: ${err}"
        }
    }

    stage('Push Image') {
        try {
            docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                app.push("${env.BUILD_NUMBER}")
            }
        } catch (err) {
            error "Failed to push Docker image: ${err}"
        }
    }

    stage('Trigger Manifest Update') {
        try {
            echo 'Triggering manifest update'
            build job: 'updatemanifest', 
                  parameters: [string(name: 'DOCKERTAG', value: "${env.BUILD_NUMBER}")]
        } catch (err) {
            error "Failed to trigger manifest update: ${err}"
        }
    }

    stage('Cleanup') {
        try {
            sh "docker rmi meshboy/godeploy:${env.BUILD_NUMBER}"
        } catch (err) {
            echo "Cleanup failed, ignoring: ${err}"
        }
    }
}