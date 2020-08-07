import axios from "axios";

const apiUrl = axios.create({
    baseURL: "/api/",
    responseType: "json"
});

const restApi = {

    // возвращает коммит гита бекенда, для девелопмента
    GitRevision: async () => {
        res = await apiUrl.get("GitRevision")
    },

    // NewGame создаёт игру, возвращает ид сессии
    NewGame: async (newGameData) => {
        res = await apiUrl.put("NewGame", newGameData )
    },

    // LoadGame грузит стейт игры по ид сессии
    LoadGame: async (sessionID) => {
        res =  await apiUrl.get("LoadGame", sessionID )
    },

    // MakeTurn делает ход и возращает новый стейт игры
    MakeTurn: async (turnData) => {
        res = await apiUrl.post("MakeTurn", turnData )
    },

    // возращает сводку по топ-10 последних игр
    TopScore: async () => {
        res = await apiUrl.get("TopScore")
    },

}

export default restApi