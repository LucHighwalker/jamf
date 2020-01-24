if [ "$ENVIRONMENT" = "dev" ]; then
    nodemon
else
    ts-node ./index.ts
fi