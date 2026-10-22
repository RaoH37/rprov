# zmprov command coverage

Status of the zmprov commands in rprov, and whether the local gosoap library
can support the missing ones.

- Last updated: 2026-10-22 (after adding `-f`)
- gosoap state: commit `47b483e` (local checkout, see `go.work`)
- The zmprov command list is based on `zmprov help all` from Zimbra 8.8/10, written
  from memory. Check it against a real `zmprov` before relying on it.

## Legend

| Column | Value | Meaning |
|---|---|---|
| rprov | ✅ | Implemented |
| | ❌ | Not implemented |
| gosoap | ✅ | The SOAP request exists in gosoap |
| | ⚠️ | No dedicated request, but the command can be built from other requests or local computation |
| | ❌ | The SOAP request is missing from gosoap |
| | — | Not applicable (already implemented) |

"ID only" means the request only accepts a zimbraId: rprov resolves the name
first with a `get<Object>Id` helper from `soap/admin/common.go`.

## Summary

| Category | Implemented | Doable with gosoap | Blocked by gosoap |
|---|---|---|---|
| Accounts | 14 | 3 (`cabulk`, `gid`, `gsig`) | 13 |
| Domains | 6 | 2 (`gdpak`, `gdpa`) | 7 |
| COS | 7 | 0 | 0 |
| Distribution lists | 11 | 0 | 1 |
| Calendar resources | 6 | 1 (`scr`, approximate) | 0 |
| Servers | 3 | 4 (`gamau`, `garpu`, `garpb`, `gamcs`) | 7 |
| Other categories | 0 | 0 | all |

## Accounts

| Command | Long name | rprov | gosoap | Notes |
|---|---|---|---|---|
| `ca` | createAccount | ✅ | — | Password given as an argument |
| `da` | deleteAccount | ✅ | — | |
| `ga` | getAccount | ✅ | — | Name or zimbraId |
| `gaa` | getAllAccounts | ✅ | — | |
| `gaaa` | getAllAdminAccounts | ✅ | — | Built on `SearchDirectoryRequest`; the admin filter must be checked against zmprov |
| `ma` | modifyAccount | ✅ | — | |
| `aaa` | addAccountAlias | ✅ | — | |
| `raa` | removeAccountAlias | ✅ | — | |
| `sp` | setPassword | ✅ | — | Password given as an argument; the policy warning in the response is not decoded |
| `ra` | renameAccount | ✅ | — | |
| `cpe` | changePrimaryEmail | ✅ | — | Chains RemoveAccountAlias, RenameAccount and AddAccountAlias; not atomic |
| `sa` | searchAccounts | ✅ | — | |
| `rmc` | recalculateMailboxCounts | ✅ | — | Multi-server behaviour to check |
| `gqu` | getQuotaUsage | ✅ | — | gosoap ignores `limit`/`offset`/`more`: results may be truncated |
| `cabulk` | createBulkAccounts | ❌ | ⚠️ | `CreateAccountRequest` in a loop; needs a password strategy and a count limit |
| `gid` | getIdentities | ❌ | ⚠️ | `DelegateAuthRequest` + `ZMailbox.GetIdentitiesRequest`; the delegated token must never be printed |
| `gsig` | getSignatures | ❌ | ⚠️ | `DelegateAuthRequest` + `ZMailbox.GetSignaturesRequest` |
| `cid` / `mid` / `did` | create/modify/deleteIdentity | ❌ | ❌ | |
| `csig` / `msig` / `dsig` | create/modify/deleteSignature | ❌ | ❌ | |
| `cds` / `mds` / `dds` / `gds` | create/modify/delete/getDataSource(s) | ❌ | ❌ | |
| `gam` | getAccountMembership | ❌ | ❌ | `GetAccountMembershipRequest` |
| `cps` | checkPasswordStrength | ❌ | ❌ | `CheckPasswordStrengthRequest` |
| `gmi` | getMailboxInfo | ❌ | ❌ | `GetMailboxInfoRequest` |

## Domains

| Command | Long name | rprov | gosoap | Notes |
|---|---|---|---|---|
| `gd` | getDomain | ✅ | — | Name or zimbraId |
| `gad` | getAllDomains | ✅ | — | |
| `cd` | createDomain | ✅ | — | |
| `md` | modifyDomain | ✅ | — | |
| `dd` | deleteDomain | ✅ | — | The server refuses non-empty domains |
| `cad` | createAliasDomain | ✅ | — | GetDomain + CreateDomain; attributes to check against zmprov; no IDN conversion |
| `gdpak` | generateDomainPreAuthKey | ❌ | ⚠️ | Key generated locally (`crypto/rand`) and stored with `ModifyDomainRequest`; the key is a secret |
| `gdpa` | generateDomainPreAuth | ❌ | ⚠️ | Reads `zimbraPreAuthKey` with `GetDomainRequest` and computes the HMAC-SHA1 locally; never print the key |
| `gdi` | getDomainInfo | ❌ | ❌ | `GetDomainInfoRequest` |
| `cta` | countAccount | ❌ | ❌ | `CountAccountRequest`; `SearchDirectoryRequest` count-only per COS would only approximate it |
| `gdsc` / `mdsc` / `rdsc` | get/modify/removeDomainSMIMEConfig | ❌ | ❌ | |
| `pfbd` | pushFreebusyDomain | ❌ | ❌ | |
| `rd` | renameDomain | ❌ | ❌ | LDAP mode only (`zmprov -l`), no SOAP equivalent |

## Classes of service (COS)

| Command | Long name | rprov | gosoap | Notes |
|---|---|---|---|---|
| `gc` | getCos | ✅ | — | Name or zimbraId |
| `gac` | getAllCos | ✅ | — | |
| `cc` | createCos | ✅ | — | Names with spaces are rejected |
| `mc` | modifyCos | ✅ | — | |
| `dc` | deleteCos | ✅ | — | Behaviour when the COS is still in use to check |
| `rc` | renameCos | ✅ | — | |
| `cpc` | copyCos | ✅ | — | gosoap does not decode `CopyCosResponse`: rprov looks the new COS up to print its ID |

## Distribution lists

| Command | Long name | rprov | gosoap | Notes |
|---|---|---|---|---|
| `gdl` | getDistributionList | ✅ | — | Name or zimbraId |
| `gadl` | getAllDistributionLists | ✅ | — | |
| `cdl` | createDistributionList | ✅ | — | |
| `cddl` | createDynamicDistributionList | ✅ | — | |
| `mdl` | modifyDistributionList | ✅ | — | |
| `ddl` | deleteDistributionList | ✅ | — | Optional `true\|false` sent as `cascadeDelete`; its exact effect to check |
| `rdl` | renameDistributionList | ✅ | — | |
| `adla` | addDistributionListAlias | ✅ | — | |
| `rdla` | removeDistributionListAlias | ✅ | — | |
| `adlm` | addDistributionListMember | ✅ | — | |
| `rdlm` | removeDistributionListMember | ✅ | — | Needs gosoap `47b483e` or later (request name typo fixed); not in the published v1.9.0 |
| `gdlm` | getDistributionListMembership | ❌ | ❌ | `GetDistributionListMembershipRequest` |

## Calendar resources

| Command | Long name | rprov | gosoap | Notes |
|---|---|---|---|---|
| `gcr` | getCalendarResource | ✅ | — | Name or zimbraId |
| `gacr` | getAllCalendarResources | ✅ | — | |
| `ccr` | createCalendarResource | ✅ | — | `displayName` and `zimbraCalResType` checked locally; password given as an argument |
| `mcr` | modifyCalendarResource | ✅ | — | |
| `dcr` | deleteCalendarResource | ✅ | — | |
| `rcr` | renameCalendarResource | ✅ | — | |
| `scr` | searchCalendarResources | ❌ | ⚠️ | No `SearchCalendarResourcesRequest`; `SearchDirectoryRequest` with type `resources` and the criteria translated to an LDAP filter would approximate it |

## Servers

| Command | Long name | rprov | gosoap | Notes |
|---|---|---|---|---|
| `gs` | getServer | ✅ | — | Name or zimbraId |
| `gas` | getAllServers | ✅ | — | |
| `ms` | modifyServer | ✅ | — | |
| `gamau` | getAllMtaAuthURLs | ❌ | ⚠️ | From `GetAllServersRequest`; selection criteria and output format to check against zmprov |
| `garpu` | getAllReverseProxyURLs | ❌ | ⚠️ | Same |
| `garpb` | getAllReverseProxyBackends | ❌ | ⚠️ | Same |
| `gamcs` | getAllMemcachedServers | ❌ | ⚠️ | From `GetAllServersRequest("memcached")` |
| `cs` | createServer | ❌ | ❌ | `CreateServerRequest` |
| `ds` | deleteServer | ❌ | ❌ | `DeleteServerRequest` |
| `gaas` | getAllActiveServers | ❌ | ❌ | `GetAllActiveServersRequest` |
| `garpd` | getAllReverseProxyDomains | ❌ | ❌ | |
| `fc` | flushCache | ❌ | ❌ | `FlushCacheRequest` |
| `rmcc` | reloadMemcachedClientConfig | ❌ | ❌ | |
| `gmcc` | getMemcachedClientConfig | ❌ | ❌ | |

## Other categories

None of these is implemented, and gosoap has none of their requests.

| Category | Commands |
|---|---|
| Global config | `gcf` (getConfig), `gacf` (getAllConfig), `mcf` (modifyConfig) |
| Mailbox and index | `sm` (selectMailbox), `ulm` (unlockMailbox), `rim` (reIndexMailbox), `cim` (compactIndexMailbox), `gis` (getIndexStats), `vi` (verifyIndex) |
| Search and GAL | `sg` (searchGal), `acg` (autoCompleteGal), `syg` (syncGal), `cto` (countObjects) |
| Rights (ACL) | `grr`, `rvr`, `ckr`, `gg`, `ger`, `gaer`, `gar`, `gr`, `grd`, `gcoa` |
| Free/busy | `gafbp`, `gfbqi`, `pfb` |
| Account loggers | `aal`, `ral`, `gal`, `gaal`, `rlog` |
| Shares | `gsi` (getShareInfo) |
| S/MIME (global) | `gcsc`, `mcsc`, `rcsc` |
| AlwaysOn clusters | `caoc`, `daoc`, `gaoc`, `gaaoc`, `maoc` |
| UC services | `cucs`, `ducs`, `gucs`, `gaucs`, `mucs`, `rucs` |
| XMPP components (removed in recent versions) | `cxc`, `dxc`, `gxc`, `gaxcs`, `mxc` |
| Calendar cache | `pacc` (purgeAccountCalendarCache) |
| zmprov shell | `desc` (describe), `exit`, interactive mode, reading commands from stdin |

## Batch mode

`rprov -f file` runs the commands of a file, one per line, like `zmprov -f`:

- a single authentication for the whole file;
- empty lines and lines starting with `#` are skipped;
- arguments can be quoted with `'` or `"` (with `\"` and `\\` escapes), `""` gives an empty value;
- a failed line is reported as `ERROR: line N: ...` (the line content is never printed, it may hold a password) and the next lines still run;
- an authentication failure stops the batch, to avoid locking the admin account with repeated attempts;
- the exit code is 1 if any line failed.

## rprov-only commands

| Command | Notes |
|---|---|
| `glc` | Prints `localconfig.xml`, like `zmlocalconfig`; secrets are hidden unless `--secrets` is given. Not a zmprov command. |

## gosoap requests not used by any zmprov command

`GetAllMailboxesRequest`, `BackupQueryRequest` (zmbackupquery), `GetLicenseRequest`
(zmlicense), `NoOpRequest`.
