#!groovy
properties([disableConcurrentBuilds(), gitLabConnection(''), pipelineTriggers([])])
def git_url
node{
    checkout scm
    git_url = scm.userRemoteConfigs[0].url
}
def Jenkinsfile = fileLoader.fromGit(env.PIPELINE_TMP,env.DEFAULT_CONFIG_REPO_V2, env.DEFAULT_CONFIG_BRANCH, 'Jenkins-Gitlab', '')
Jenkinsfile.start_pipeline(git_url)