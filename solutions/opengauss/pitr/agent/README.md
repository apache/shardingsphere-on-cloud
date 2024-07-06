# PITR/AGENT

PITR agent server for openGauss.

## project description

### layout

* `internal`:For internal code, use golang internal directory to isolate references.
	- `cons`: Constant directory.
    - `handler`: HTTP handler.
    - `pkg`: Business logic related packages.
* `pkg`: Business logic independent packages.

