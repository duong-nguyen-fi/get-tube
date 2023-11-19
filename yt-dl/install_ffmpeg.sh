dnf install -y epel-release
dnf config-manager --set-enabled crb
dnf install --nogpgcheck https://mirrors.rpmfusion.org/free/el/rpmfusion-free-release-$(rpm -E %rhel).noarch.rpm -y
dnf install --nogpgcheck https://mirrors.rpmfusion.org/nonfree/el/rpmfusion-nonfree-release-$(rpm -E %rhel).noarch.rpm -y
dnf install -y ffmpeg ffmpeg-devel