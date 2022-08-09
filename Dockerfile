FROM scratch
COPY getipinfo /usr/local/bin/getipinfo
ENTRYPOINT [ "/usr/local/bin/getipinfo" ]