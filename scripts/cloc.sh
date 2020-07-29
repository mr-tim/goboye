#!/bin/bash

cloc --fullpath --not-match-d="(node_modules)" --not-match-f="(yarn\.lock|package\.json|package\-lock\.json)" .
