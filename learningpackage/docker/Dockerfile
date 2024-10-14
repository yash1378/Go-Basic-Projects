# FROM ubuntu:20.04
# LABEL authors="YashDugriyal <yash1066@gmail.com>"
# LABEL version="1.0"
# LABEL description="This is the description of the image"
# ARG APP_PORT=8080
# RUN echo "Building app on port $APP_PORT"


# LABEL authors="YashDugriyal <yash1066@gmail.com>" version="1.0" description="This is the description of the image for the single line label"

# ENV MYSQL_PASSWD=test1234
# RUN apt-get update && apt-get install -y nginx
# EXPOSE 100
# # RUN apt-get update
# RUN ["apt-get","-y","upgrade"]
# COPY index.html /usr/share/nginx/html/
# COPY index.html /var/www/html/
# RUN /bin/bash -c "apt-get -y dist-upgrade"

# ENTRYPOINT ["kiski"]
# CMD ["Hello World"]
# CMD ["nginx","-g","daemon off;"]

# Start with a base image
FROM ubuntu:latest

# Use /bin/sh (default) to run a simple command
RUN echo "This is running with /bin/sh" > /sh_output.txt

# Change to /bin/bash for more complex commands
SHELL ["/bin/bash", "-c"]

# Use bash to run a complex command
RUN echo "This is running with /bin/bash" > /bash_output.txt && \
    echo "Bash allows us to use && for chaining commands" >> /bash_output.txt

# Check the content of the files to confirm both commands executed
# RUN cat /sh_output.txt /bash_output.txt
