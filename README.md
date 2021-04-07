Stakater's OpenShift oauth-proxy
=====================

This is a fork of the https://github.com/openshift/oauth-proxy project with change in the sign_in and error
page customized for the use of Stakater App Agility Platform. 

## Stakater's Customizations

- Sign in page
- Errors page
- alpine:3.13 as base image image

The changed templates can be found under stakater/templates dir. These templates have been hardcoded into `templates.go` file

Plz find detailed README here: https://github.com/openshift/oauth-proxy