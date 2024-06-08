FROM --platform=$TARGETPLATFORM alpine:latest
ARG TARGETPLATFORM

COPY ./dist/$TARGETPLATFORM/* /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/the_goal_game" ]
CMD [ "-plot-achieved-output" ]