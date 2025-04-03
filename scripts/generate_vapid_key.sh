#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
ENV_FILE="${PROJECT_DIR}/.env"

echo "VAPID_KEYがすでに存在するか確認しています..."
if [ -f "${ENV_FILE}" ]; then
    if grep -q "VAPID_PRIVATE_KEY" "${ENV_FILE}"; then
        echo "[FAILED] VAPID_KEYはすでに存在します。"
        exit 0
    fi
else
    echo "[FAILED] .envファイルが存在しません。"
    exit 1
fi

echo "VAPID_KEYを生成します..."
openssl ecparam -genkey -name prime256v1 -noout -out soruceKey.pem
VAPID_PRIVATE_KEY=$(openssl pkcs8 -in soruceKey.pem -topk8 -nocrypt -outform DER | base64 | tr -d '\n')
VAPID_PUBLIC_KEY=$(openssl ec -in soruceKey.pem -pubout -outform DER | base64 | tr -d '\n')

echo "\nVAPID_PRIVATE_KEY=${VAPID_PRIVATE_KEY}" >> "${ENV_FILE}"
echo "VAPID_PUBLIC_KEY=${VAPID_PUBLIC_KEY}" >> "${ENV_FILE}"

echo "[SUCCESS] VAPID_PRIVATE_KEYとVAPID_PUBLIC_KEYが.envに追加されました。"

rm soruceKey.pem
echo "soruceKey.pemが削除されました。"
