# 1. Configuration


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

# 2. Usage

`assume-role dev`