import { useMainGetProfileHook } from "../hooks"

export const MainProfile = () => {
  const { mainImg } = useMainGetProfileHook()

  return (
    <div className = "mainProfileContainer">
      
      {/* 메인 이미지로 쓰이는 컴퍼넌트 */}
      <div className = "mainProfileImgBox">
        {
          mainImg ? 
          <img src={mainImg} alt="메인이미지" className = "mainProfileImg" /> : "로딩 중"
        }
      </div>

      {/* 메인 화면에 사용될 글 */}
      <div className = "mainProfileText">
        <h1 className = "mainProfileTextValue">
          동완 ♡ 서희
        </h1>
      </div>

    </div>
  )
}
