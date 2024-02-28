FROM scratch

LABEL maintainer="Ben Sandberg <info@pdxfixit.com>" \
      name="hostdb-collector-oneview" \
      vendor="PDXfixIT, LLC"

COPY hostdb-collector-oneview /usr/bin/
COPY config.yaml /etc/hostdb-collector-oneview/

ENTRYPOINT [ "/usr/bin/hostdb-collector-oneview" ]
