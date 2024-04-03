# 1. Install
## 1.1. Linux/Macos amd64
```shell
sudo curl https://github.com/lttrungn1742/assume-role/releases/download/v1/assume-role --output /usr/local/bin/assume-role
```

# 2. Configuration


For example:

- `~/.aws/config`
```
[default]
region = ap-southeast-1
output = json

[profile dev]
region = ap-southeast-1
role_arn = arn:aws:iam::************:role/Admin
mfa_serial = arn:aws:iam::************:mfa/Authy
source_profile = default

[profile prod]
region = ap-southeast-1
role_arn = arn:aws:iam::************:role/Admin
mfa_serial = arn:aws:iam::************:mfa/Authy
source_profile = default
```

- `~/.aws/credentials`
```
[default]
aws_access_key_id = ****************
aws_secret_access_key = ****************
```

# 3. Usage

`assume-role dev`