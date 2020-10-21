%define install_dir %{buildroot}/usr/local/gohttpd
%define systemd_dir %{buildroot}/usr/lib/systemd/system

%define package_name gohttpd
%define package_conf %{install_dir}/%{package_name}.conf

Name:           %{package_name}
Summary:        Simple HTTP Server
Version:        1.1.0
Release:        1013%{?dist}
Group:          System Environment/Daemons
License:        GPLv2
URL:            https://github.com/yodstar
Packager:       baoqiang <yodstar@foxmail.com>
Vendor:         yodstar
Source0:        %{name}-%{version}.tar.gz
#patch0:        a.patch
BuildRoot:      %_topdir/BUILDROOT         
#BuildRequires:  go,make
#Requires:      pcre,pcre-devel,openssl,chkconfig  #软件运行需要的软件包，也可以指定最低版本如 bash >= 1.1.1

%description
Simple HTTP Server

%prep
%setup -q
# %patch0 -p1
# %%configure

%build
go build -o %{package_name} -ldflags "-w -s" %{package_name}

# %pre
if [ -f %{package_conf} ];then mv %{package_conf} %{package_conf}.old;fi

%install
rm -rf %{buildroot}
mkdir -p %{systemd_dir}
install -d -m0755 %{install_dir}/html
install -m0755 ./%{package_name} %{install_dir}/%{package_name}
install -m0600 ./%{package_name}.conf %{package_conf}
install -m0755 ./%{package_name}.service %{systemd_dir}/%{package_name}.service
install -m0644 ./html/favicon.ico %{systemd_dir}/html/favicon.ico
install -m0644 ./html/index.html %{systemd_dir}/html/index.html

%clean
rm -rf %{buildroot}

%post
%systemd_post %{package_name}.service

%files  
%defattr (-,root,root)
/usr/local/gohttpd
/usr/lib/systemd/system

%postun
%systemd_postun_with_restart %{package_name}.service

#%preun
%systemd_preun %{package_name}.service

%changelog 
* Tue Oct 21 2020 baoqiang <yodstar@foxmail.com> - 1.1.0
- Release version