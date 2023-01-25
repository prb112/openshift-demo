# How to use OpenScap Scanner on a Mac

For those, not yet using `openscap-scanner` on their systems, [OpenSCAP](https://www.open-scap.org/) is an security auditing framework that utilizes the Extensible Configuration Checklist Description Format (XCCDF) and the `openscap-scanner` executes over the security *profile* on a target system. 

One gotcha, I have a Mac, and the tool is not natively supported on the Mac. I decided to use it through a `fedora` container running in [Podman](https://podman.io/). 

Here are the steps to running on a Mac with [`complianceascode/content`](https://github.com/ComplianceAsCode/content/releases/tag/v0.1.65)'s release.

# Steps

1. Download the Docker File 

<< Add notes about the DockerFile>>

2. Build the Image

```
$ podman build -f Dockerfile -t ocp-power.xyz/compliance/openscap-wrapper:latest
...
```

3. Download the content files [scap-security-guide-0.1.65.zip](https://github.com/ComplianceAsCode/content/releases/download/v0.1.65/scap-security-guide-0.1.65.zip)

```
$ curl -O -L https://github.com/ComplianceAsCode/content/releases/download/v0.1.65/scap-security-guide-0.1.65.zip
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100  130M  100  130M    0     0  2752k      0  0:00:48  0:00:48 --:--:-- 5949k
```

4. Unzip the `scap-security-guide-0.1.65.zip` file.

```
$ unzip scap-security-guide-0.1.65.zip
```

5. Rename the directory `scap-security-guide-0.1.65` to `scap`

```
$ mv scap-security-guide-0.1.65 scap
```

6. List the profiles in a specific XML. 

```
$ podman run --rm -v ./scap:/scap ocp-power.xyz/compliance/openscap-wrapper:latest oscap info --profiles /scap/ssg-ocp4-ds.xml
xccdf_org.ssgproject.content_profile_cis-node:CIS Red Hat OpenShift Container Platform 4 Benchmark
xccdf_org.ssgproject.content_profile_cis:CIS Red Hat OpenShift Container Platform 4 Benchmark
xccdf_org.ssgproject.content_profile_e8:Australian Cyber Security Centre (ACSC) Essential Eight
xccdf_org.ssgproject.content_profile_high-node:NIST 800-53 High-Impact Baseline for Red Hat OpenShift - Node level
xccdf_org.ssgproject.content_profile_high:NIST 800-53 High-Impact Baseline for Red Hat OpenShift - Platform level
xccdf_org.ssgproject.content_profile_moderate-node:NIST 800-53 Moderate-Impact Baseline for Red Hat OpenShift - Node level
xccdf_org.ssgproject.content_profile_moderate:NIST 800-53 Moderate-Impact Baseline for Red Hat OpenShift - Platform level
xccdf_org.ssgproject.content_profile_nerc-cip-node:North American Electric Reliability Corporation (NERC) Critical Infrastructure Protection (CIP) cybersecurity standards profile for the  Red Hat OpenShift Container Platform - Node level
xccdf_org.ssgproject.content_profile_nerc-cip:North American Electric Reliability Corporation (NERC) Critical Infrastructure Protection (CIP) cybersecurity standards profile for the  Red Hat OpenShift Container Platform - Platform level
xccdf_org.ssgproject.content_profile_pci-dss-node:PCI-DSS v3.2.1 Control Baseline for Red Hat OpenShift Container Platform 4
xccdf_org.ssgproject.content_profile_pci-dss:PCI-DSS v3.2.1 Control Baseline for Red Hat OpenShift Container Platform 4
```

7. Details on the profile

```
$ podman run --rm  -v ./scap:/scap ocp-power.xyz/compliance/openscap-wrapper:latest oscap info --profile xccdf_org.ssgproject.content_profile_cis-node /scap/ssg-ocp4-ds.xml
Document type: Source Data Stream
Imported: 2022-12-02T19:09:36

Stream: scap_org.open-scap_datastream_from_xccdf_ssg-ocp4-xccdf.xml
Generated: (null)
Version: 1.3
Profile
        Title: CIS Red Hat OpenShift Container Platform 4 Benchmark
        Id: xccdf_org.ssgproject.content_profile_cis-node

        Description: This profile defines a baseline that aligns to the Center for Internet Security® Red Hat OpenShift Container Platform 4 Benchmark™, V1.1.  This profile includes Center for Internet Security® Red Hat OpenShift Container Platform 4 CIS Benchmarks™ content.  Note that this part of the profile is meant to run on the Operating System that Red Hat OpenShift Container Platform 4 runs on top of.  This profile is applicable to OpenShift versions 4.6 and greater.
```

8. Now, I can run more advanced commands on the profiles on my Mac.

```
$ podman run --rm  -v ./scap:/scap ocp-power.xyz/compliance/openscap-wrapper:latest oscap oval generate report /scap/ssg-ocp4-ds.xml 2>&1
```

## References
1. [OpenScap Downloads](https://www.open-scap.org/download/)
2. [OpenScap source code](https://github.com/OpenSCAP/openscap)
3. [OpenScap Manual Source](https://github.com/OpenSCAP/openscap/blob/maint-1.3/docs/manual/manual.adoc)
4. [OpenScap Manual Published](https://static.open-scap.org/openscap-1.2/oscap_user_manual.html#_installation)

## Notes
Note, I found I had to do the following on my Mac to get the volume to mount. 

```
$ podman machine stop
$ podman machine set --rootful
$ podman machine start
$ sudo /opt/homebrew/Cellar/podman/4.3.1/bin/podman-mac-helper install
$ podman machine stop; podman machine start
```