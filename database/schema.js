// users

{
    "email": "jet@any.tv",
    "alias": "craeden",
    "roles": ['dm', 'c', 'gm'],
    "created_at": +Date,
    "updated_at": null
}

//  projects/dungeons

{
    "name": "New Dashboard",
    "description": "Freedom! Dashboard rewrite",
    "tech_stack": ["php", "mysql"],
    "dungeon_master": [ObjectId(), ObjectId(), ...], // user._id
    "guardians": [ObjectId(), ObjectId(), ...], // user._id
    "created_at": +Date,
    "updated_at": null
}

// tasks/quests

{
    "name": "Feature magic",
    "dungeon_id": ObjectId(),
    "description": "this feature is so awesome...",
    "checklist": [
        {
            "description": "",
            "status": "done"
        },
        {
            "description": "",
            "status": "done"
        }
    ],
    "members": [ ObjectId(), ObjectId(), ...], // user._id
    "class": "S",
    "start_date": Date?,
    "due_date": Date?
    "created_at": +Date,
    "updated_at": null
}

// thread-entry

{
    "target_id": ObjectId,  // quest._id or dungeon._id
    "entry": "this shit is potato",
    "history": [],
    "created_at": +Date,
    "updated_at": null
}
