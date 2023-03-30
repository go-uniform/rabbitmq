FROM scratch
ADD build /service
ENTRYPOINT ["/service"]
HEALTHCHECK CMD --interval=2s --timeout=10s --start-period=10s service command:ping
CMD ["run"]