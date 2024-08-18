#!/bin/bash

name="Nobuchika Tanaka"
email="n.tanaka@idcf.jp"

sh -c "$(curl -fsLS get.chezmoi.io)" -- \
   -b "${HOME}/.local/bin" \
   -d init --apply n-tanaka-idcf \
           --promptString name="${name}" \
           --promptString email="${email}"
