# Kofini

Kofini means "basket" in Greek. This project is a CLI tool that will help users
discover and transform their, initially targeting PCF applications to OCP.

## Description

This effort will cover the development of a golang CLI tool capable of
performing discovery and transformation of a PCF application. This task does not
expect a tool compatible with the Konveyor addon process.

We want a `discover` and `generate` command. The `discover` command should
output a yaml/json file similar to the output of M2K's `collect` and `plan`
commands. The format of this file can be a bare minimum for an MVP; though we
should consider that we want this to eventually be a canonical configuration
structure capable of handling other platforms. This `discover` command can take
in one supported `–platform` flag: `pcf`.

The `generate` command should template out assets for the application to be
deployed to OCP. This `generate` command should take in a helm template and spit
out a helm chart for a user to use for the asset deployment.
