#!/bin/bash
# Update AppEngine app: get deps, update, rm deps.

main() {
  cd "$(dirname "${BASH_SOURCE[0]}")"

  # Add dependencies
  mkdir -p github.com/StalkR || err "failed mkdir github.com/StalkR"
  git clone https://github.com/StalkR/imdb.git github.com/StalkR/imdb || \
    err "failed git clone"

  # Update
  HOME=. appcfg.py --noauth_local_webserver --oauth2 update .

  # Cleanup
  rm -rf github.com
}

err() {
  echo "$@" >&2
  exit 1
}

if [[ "${BASH_SOURCE[0]}" = "$0" ]]; then
  main "$@"
fi
