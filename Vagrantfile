Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"

  config.vm.define 'noaa-env-centos' do |config|
    config.vm.hostname = 'noaa-env-centos.local'
    config.vm.network "forwarded_port", guest: 8080, host: 8080, guest_ip: "127.0.0.1", auto_correct: true
    config.vm.synced_folder ".", "/vagrant", type: "nfs", mount_options: ["dmode=777", "fmode=776"]
    config.vm.provision "shell", path: "provision.sh"
  end
end