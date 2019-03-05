FROM envoyproxy/envoy:v1.9.0
COPY envoy.yaml /etc/envoy.yaml
COPY ./keys /keys
CMD /usr/local/bin/envoy -c /etc/envoy.yaml
