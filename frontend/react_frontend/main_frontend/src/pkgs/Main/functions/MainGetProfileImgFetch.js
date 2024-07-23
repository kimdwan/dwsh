

export const MainGetProfileImgFetch = async ( url ) => {
  try {
    const response = await fetch(url, {
      method : "GET",
      headers : {
        "Content-Type" : "application/json; charset=utf-8",
        "X-Requested-With": "XMLHttpRequest",
      },
      credentials : "include",
    })
    
    if (!response.ok) {
      if (response.status === 500) {
        throw new Error("서버에서 이미지를 불러오는 중 오류가 발생했습니다")
      } else {
        throw new Error(`오류가 발생했습니다 오류번호: ${response.status}`)
      }
    }

    const data = await response.json()

    return data

  } catch (err) {
    throw err
  }
}