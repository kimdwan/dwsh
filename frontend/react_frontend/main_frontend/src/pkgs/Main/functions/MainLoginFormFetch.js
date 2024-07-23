

export const MainLoginFormFetch = async ( url, datas, setEmailErrorMsg, setPasswordErrorMsg ) => {

  try {
    const response = await fetch(url, {
      method : "POST",
      headers : {
        "Content-Type" : "application/json",
        "X-Requested-With": "XMLHttpRequest",
      },
      body : JSON.stringify(datas),
      credentials : "include"
    })

    setEmailErrorMsg(null)
    setPasswordErrorMsg(null)

    if (!response.ok) {

      if (response.status === 404) {
        setEmailErrorMsg("이메일이 존재하지 않습니다")
        throw new Error("이메일이 존재하지 않습니다")
      } else if (response.status === 401) {
        setPasswordErrorMsg("비밀번호가 틀렸습니다")
        throw new Error("비밀번호가 틀렸습니다")
      } else if (response.status === 500) {
        throw new Error("서버에 오류가 발생했습니다")
      } else {
        alert(`오류발생: ${response.status}`)
        throw new Error(`오류가 발생했습니다 오류 번호: ${response.status}`)
      }
    }

    const backend_data = await response.json()

    return backend_data

  } catch (err) {
    throw err
  }
}