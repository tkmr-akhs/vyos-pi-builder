FROM vyos/vyos-build:current-arm64

#RUN mv /etc/apt/sources.list /etc/apt/sources.list.orig

RUN echo "deb http://deb.debian.org/debian bookworm main contrib non-free" >> /etc/apt/sources.list
RUN echo "deb-src http://deb.debian.org/debian bookworm main contrib non-free" >> /etc/apt/sources.list

RUN echo "deb http://deb.debian.org/debian bookworm-updates main contrib non-free" >> /etc/apt/sources.list
RUN echo "deb-src http://deb.debian.org/debian bookworm-updates main contrib non-free" >> /etc/apt/sources.list

RUN echo "deb http://deb.debian.org/debian bookworm-backports main contrib non-free" >> /etc/apt/sources.list
RUN echo "deb-src http://deb.debian.org/debian bookworm-backports main contrib non-free" >> /etc/apt/sources.list

RUN echo "deb http://security.debian.org/debian-security/ bookworm-security main contrib non-free" >> /etc/apt/sources.list
RUN echo "deb-src http://security.debian.org/debian-security/ bookworm-security main contrib non-free" >> /etc/apt/sources.list

RUN apt -y install gdisk
RUN apt -y install kpartx

RUN rm -f /etc/apt/sources.list
#RUN mv /etc/apt/sources.list.orig /etc/apt/sources.list

