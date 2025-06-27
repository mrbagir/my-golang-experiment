# Fast Port Forward

## How to Run

### 1. Edit Configuration (`forwarder.ini`)

```ini
[Config]
print_config=true          # Display configuration before running
remotePort=9090            # Default gRPC port in Kubernetes/OpenShift
kubeConfigPath="~/.kube/config" # Default kubeconfig path

[PortForward.dev]          # .dev/.prestage/.stage/.prerelease/.release
namespace=default          # Project name / Namepace
addons-auth-service=9105   # [service_name]=[local_port]
addons-task-service=9090
```

### 2. Login to OpenShift

1. Open your terminal.
2. Run the following command:

```bash
# Usage:
# oc login <server_url> -u <username> -p <password> --insecure-skip-tls-verify

oc login https://example.com:6443 -u admin -p admin123@! --insecure-skip-tls-verify
```

### 3. Run Forwarder

```bash
# Usage:
# forwarder [dev|prestage|stage|prerelease|release]

# Default to 'dev'
make run

# Run for a specific environment (e.g., dev)
make run dev
```

### 4. Add to PATH (Optional)

To make the `forwarder` command available globally, add its directory to your system `PATH`.

---

## Screenshot

![image.png](docs\image.png)