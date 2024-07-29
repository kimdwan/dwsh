import { useState, useEffect } from "react"
import { MainGetProfileImgFetch } from "../functions"

export const useMainGetProfileHook = () => {
  const [mainImg, setMainImg] = useState("")
  const go_backend_url = process.env.REACT_APP_GO_BACKEND_URL

  useEffect(() => {
    const url = `${go_backend_url}/etc/main/profile`

    const getMainProfileFetch = async (url) => {
      try {
        const response = await MainGetProfileImgFetch(url)
        if (response) {
          setMainImg(`data:image/${response["imgType"]};base64,${response["base64Img"]}`)
        }

      }catch (err) {
        alert(err)
        throw err
      }
    }
    getMainProfileFetch(url)
  }, [go_backend_url])

  return { mainImg }
}