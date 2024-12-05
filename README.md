# GDG_backend_sideProject
# gRPC를 활용한 비트코인 시세 모니터링 서버 및 GCP를 통한 관리/배포 시스템
  
### 📍배포 링크:  [http://35.216.21.2:8080/](http://35.216.21.2:8080/)
## Contributors ✨
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
---

### ✅ 프로젝트 개요

```go
본 프로젝트는 GCP 환경에서 gRPC api를 활용해 Bitcoin 정보를 가져와
차트 형태로 보여주며 코인 정보 모니터링을 도와주는 프로젝트입니다.
```

### ✅ 아키텍처 키워드

```go
- gRPC: 각 서비스는 gRPC로 통신
- GCP : 구글 클라우드 환경에 적용
```

### ✅ 주요 기능
- **REST API를 지원하는 gRPC 서버 구축** : 효율적이고 확장 가능한 서비스를 위한 gRPC 기반 서버 제공
- **CandleStick chart** : 비트코인 데이터 시각화
- **GCP 서버 배포** : Google Compute Engine(GCE)로 배포
----
![서비스 화면](https://github.com/user-attachments/assets/7c727101-1671-4379-96d3-17513c067777)

---
프로젝트 구조
```
GDG_backend_sideProject/
├── proto/                          
│   └── monitoring.proto             # gRPC 메서드와 메시지 구조 정의
├── server/                          
│   ├── api_server/                  # API 서버 디렉토리
│   │   ├── tmp/                     
│   │   ├── Dockerfile              
│   │   ├── go.mod                   
│   │   ├── go-1.21.1.linux-amd64.tar.gz # Go 실행 파일 압축 패키지
│   │   └── main.go                  # API 서버의 진입점, 서버 실행 코드
│   ├── chart_server/                # 차트 서버 디렉토리
│   │   ├── chart/                   
│   │   │   └── kline_chart.go       # 캔들 차트 생성 및 데이터 관리
│   │   ├── grpc/                    
│   │   │   └── grpc_client.go       # gRPC 클라이언트, 실시간 데이터 수신
│   │   ├── handler/                
│   │   │   └── chart_handler.go     # HTTP 요청 처리 및 차트 렌더링
│   │   ├── model/                   
│   │   │   └── chart_data.go        # 차트 데이터 관리 및 갱신
│   │   ├── tmp/                     
│   │   ├── util/                   
│   │   │   └── data_processor.go    # 거래 데이터 그룹화 및 분석
│   │   ├── Dockerfile              
│   │   ├── go.mod                   
│   │   ├── go-1.21.1.linux-amd64.tar.gz 
│   │   └── main.go                  # 차트 서버의 진입점, 서버 실행 코드
│   ├── pb/                         
│   │   ├── monitoring.pb.go         # gRPC Stub 및 메시지 구조 정의
│   │   └── monitoring_grpc.pb.go    # gRPC 서비스 Stub 구현
│   └── tmp/                         # 서버에서 사용하는 임시 디렉토리
│       └── build-errors.log         # 빌드 및 실행 중 발생한 에러 로그
├── docker-compose.yaml             
├── go.mod                           
└── Makefile                         

```

