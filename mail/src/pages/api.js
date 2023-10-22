import axios from "axios";

export async function fetchAllUsers (){
    const { data } = await axios.get('/users')
    // console.log('data',data.data);
    
 
    console.log('data',data.users);

    return  data.users
}

export async function fetchAllMails (){
    const { data } = await axios.get('/mails')
    // console.log('data',data.data);
    
 
    console.log('data',data);

    return data.mails
}

export async function fetchAllCrafts (){
    const { data } = await axios.get('/crafts')
    // console.log('data',data.data);
    
 
    console.log('data',data);

    return data.crafts
}

export async function fetchACraft (id){
    console.log('id',id);
    
    const { data } = await axios.get('/crafts/'+id)
    // console.log('data',data.data);
 
    console.log('craft',data.craft);

    return data.craft
}