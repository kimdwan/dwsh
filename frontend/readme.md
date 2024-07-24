# 동완 서희 어플리케이션 프론트엔드 서비스

# 사용 프론트엔드 
1. react (typescript)

- 사실상 이걸로 대부분의 서비스를 구현한다.

# 배운점 (프론트엔드에서도 꾸준히 배운다)
1. npx create-react-app filename --template typescript
-> typescript 버전의 react app이 생성된다.
2. content-type을 확인하자 
-> 백엔드에서 보낸 파일의 타입을 파악해서 이에 맞는 파싱으로 해주어야 한다
-> 기존에 에서 utf-8의 형식으로 보낸 데이터는 이러한 형식으로 파싱 "application/json; charset=utf-8"
3. img 태그가 사용할 수 있는 방식은 다음과 같다 (img 태그에 이 파일이 어떤 형식인지 알려주어야 한다.)
-> `data:(데이터 타입: 예시는 이미지)/(file형식: 예시는 jpg, jpeg, webp, ..., 등등);(인코딩 방식 예시: base64),(인코딩)`
4. useNavigate는 최상단에 위치 해야 한다. 
5. textare를 읽기 전용으로 하려면 readonly를 추가해야 한다.
6. useState의 set함수는 return 값에 따라서 값을 변경시킬수 있다.