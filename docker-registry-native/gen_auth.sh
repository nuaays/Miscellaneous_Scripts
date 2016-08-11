

mkdir -p auth && docker run --entrypoint htpasswd registry:2.5 -Bbn testuser testpassword > auth/htpasswd
