FROM rockylinux:9.2
RUN yum update -y
RUN yum upgrade -y python3; \
    yum install -y python3-pip;
RUN python3 -m pip install -U yt-dlp
COPY install_ffmpeg.sh ./install_ffmpeg.sh
RUN /bin/bash -c 'chmod a+x ./install_ffmpeg.sh; ./install_ffmpeg.sh'
RUN useradd -ms /bin/bash cloud-user
USER cloud-user
RUN mkdir -p /files
RUN chown cloud-user:cloud-user /files
WORKDIR /files
ENTRYPOINT ["yt-dlp"]
