# GDG_backend_sideProject
# gRPC를 활용한 비트코인 시세 모니터링 서버 및 GCP를 통한 관리/배포 시스템
  
📍 **배포 링크**:  
[**http://35.216.21.2:8080/**](http://35.216.21.2:8080/)

### 💡 주요 기능
- 🧩 **REST API를 지원하는 gRPC 서버 구축** : 효율적이고 확장 가능한 서비스를 위한 gRPC 기반 서버 제공
- 📊 **CandleStick chart** : 비트코인 데이터 시각화
- 🔒 **GCP 서버 배포** : Google Compute Engine(GCE)로 배포 

### 1️⃣ 프로젝트 개요

```go
본 프로젝트는 GCP 환경에서 gRPC api를 활용해 Bitcoin 정보를 가져와
차트 형태로 보여주며 코인 정보 모니터링을 도와주는 프로젝트입니다.
```

### 2️⃣ 아키텍처 키워드

```go
- gRPC: 각 서비스는 gRPC로 통신
- GCP : 구글 클라우드 환경에 적용
```

### 3️⃣ 기대 효과

```go
1. golang과 gRPC를 결합해 Bitcoin 스트리밍 서비스를 구현합니다.
2. gRPC를 사용해 변경된 데이터만 송수신하여 더 빠르게 변경사항을 적용합니다
3. GCP를 사용해 실시간으로 기능을 관리하고 배포를 자동화합니다. 
```

### 4️⃣ 핵심 기능

```go
1. CoinGecko api(RESTful api)를 gRPC 형태로 변환
2. 변환된 데이터를 server Streaming RPC 형태로 수신
3. 수신된 데이터를 Candlestick Chart 형태로 표현
4. 해당 서버를 gcp상에서 구현
```

## Contributors ✨

Thanks goes to these wonderful people:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/jihyeonAnAn"><img src="https://avatars.githubusercontent.com/u/84323575?v=4?s=100" width="100px;" alt="Jihyeon An"/><br /><sub><b>Jihyeon An</b></sub></a><br />
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Chocobone"><img src="https://avatars.githubusercontent.com/u/73699062?v=4?s=100" width="100px;" alt="Chocobone"/><br /><sub><b>유현웅</b></sub></a><br />
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
