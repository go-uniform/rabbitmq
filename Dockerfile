FROM scratch
ADD build /service
ENTRYPOINT ["/service"]
HEALTHCHECK CMD ["command:ping"]
CMD ["run"]