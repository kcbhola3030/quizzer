const  importData = require("./data.json");
let port = process.env.port || 5000;


const express = require("express");
// const mongoose  = require('mongoose');
// const { findOne } = require("domutils");

// mongoose.set('strictQuery', true);

// try {
//     const connect = mongoose.connect('mongodb://localhost:27017/golang',)
//     console.log("db working");
    
// } catch (error) {
//     console.log(error);
// // }
// const User = mongoose.model('User', { name: String, score: Number });



const app = express();
app.use(express.json());


app.get("/" , (req,res) => {
    res.send("Hello world");
});

app.get("/questions", (req,res) => {
    res.send(importData);
});


// app.post("/user",async(req,res)=>{
//     const {name, score} = req.body;
//     try {
//         const newUser  = new User({
    
//             name:name,
//             score:score
//         })
//         await newUser.save();
//         console.log({_id:newUser.id});
//         res.status(200).json({_id : newUser.id});
//         return
//     } catch (error) {
//         res.status(400).send(error);
//     }
// });
// //add score
// app.put('/updateUser',async(req,res)=>{
//     try {
//         const {userId, score} =req.body;
//         console.log(userId);
//         const updateUser = await User.findByIdAndUpdate(userId,{score:score});
//         res.status(200).json(updateUser);
//         return
        
//     } catch (error) {
//         console.log(error)
//         res.status(401).json(error)
//     }
// })


// app.get('/data',(req,res)=>{

// })
app.listen(port, () => {
    console.log(`Eg app is listening on port http://localhost:3000`);
})