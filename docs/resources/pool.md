---
page_title: "quortex_pool Resource - terraform-provider-quortex"
subcategory: ""
description: |-
---

# quortex_pool

Manage a pool resource. For more information see
[the official documentation](https://help.quortex.io/en/)
and [the API reference](https://web.quortex.io/documentation/ott).

## Example Usage - basic configuration

```hcl
resource "quortex_pool" "my_pool" {
  name                = "my_pool"
  streaming_countries = ["FRA"]
  input_region        = "ireland"
  published           = true
}
```

## Argument reference

The following arguments are required:

- `name` - (Required) This is the name of the pool. It can be whatever that identify the pool.
- `input_region` - (Required) This is the region on which the streams will acquired.
- `streaming_countries` - (Required) Items Enum: `ABW` `AFG` `AGO` `AIA` `ALA` `ALB` `AND` `ARE` `ARG` `ARM` `ASM` `ATA` `ATF` `ATG` `AUS` `AUT` `AZE` `BDI` `BEL` `BEN` `BES` `BFA` `BGD` `BGR` `BHR` `BHS` `BIH` `BLM` `BLR` `BLZ` `BMU` `BOL` `BRA` `BRB` `BRN` `BTN` `BVT` `BWA` `CAF` `CAN` `CCK` `CHE` `CHL` `CHN` `CIV` `CMR` `COD` `COG` `COK` `COL` `COM` `CPV` `CRI` `CUB` `CUW` `CXR` `CYM` `CYP` `CZE` `DEU` `DJI` `DMA` `DNK` `DOM` `DZA` `ECU` `EGY` `ERI` `ESH` `ESP` `EST` `ETH` `FIN` `FJI` `FLK` `FRA` `FRO` `FSM` `GAB` `GBR` `GEO` `GGY` `GHA` `GIB` `GIN` `GLP` `GMB` `GNB` `GNQ` `GRC` `GRD` `GRL` `GTM` `GUF` `GUM` `GUY` `HKG` `HMD` `HND` `HRV` `HTI` `HUN` `IDN` `IMN` `IND` `IOT` `IRL` `IRN` `IRQ` `ISL` `ISR` `ITA` `JAM` `JEY` `JOR` `JPN` `KAZ` `KEN` `KGZ` `KHM` `KIR` `KNA` `KOR` `KWT` `LAO` `LBN` `LBR` `LBY` `LCA` `LIE` `LKA` `LSO` `LTU` `LUX` `LVA` `MAC` `MAF` `MAR` `MCO` `MDA` `MDG` `MDV` `MEX` `MHL` `MKD` `MLI` `MLT` `MMR` `MNE` `MNG` `MNP` `MOZ` `MRT` `MSR` `MTQ` `MUS` `MWI` `MYS` `MYT` `NAM` `NCL` `NER` `NFK` `NGA` `NIC` `NIU` `NLD` `NOR` `NPL` `NRU` `NZL` `OMN` `PAK` `PAN` `PCN` `PER` `PHL` `PLW` `PNG` `POL` `PRI` `PRK` `PRT` `PRY` `PSE` `PYF` `QAT` `REU` `ROU` `RUS` `RWA` `SAU` `SDN` `SEN` `SGP` `SGS` `SHN` `SJM` `SLB` `SLE` `SLV` `SMR` `SOM` `SPM` `SRB` `SSD` `STP` `SUR` `SVK` `SVN` `SWE` `SWZ` `SXM` `SYC` `SYR` `TCA` `TCD` `TGO` `THA` `TJK` `TKL` `TKM` `TLS` `TON` `TTO` `TUN` `TUR` `TUV` `TWN` `TZA` `UGA` `UKR` `UMI` `URY` `USA` `UZB` `VAT` `VCT` `VEN` `VGB` `VIR` `VNM` `VUT` `WLF` `WSM` `YEM` `ZAF` `ZMB` `ZWE`
  List of countries (from ISO_3166-1_alpha-3) on which the streams will be available.

The following arguments are optional:

- `published` - (Optional) Allow the enabling/disabling of all publishing points of the pool.

## Attributes reference

- `id` - This is the universal unique identifier of the ressource.
