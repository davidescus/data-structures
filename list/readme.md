# Usage
    
Init list `{ nil }`

    l:= list.New()

Add a value `{ 1 }`

    _ = l.PushBack(1)

Add another value `{ 1 -> 2 }`
    
    middleNode := l.PushBack(2)

Add another value in front `{ "Add in front" -> 1 -> 2 }`

    frontNode := l.PushFront("Add in front")

Add another value at the end `{ "Add in front" -> 1 -> 2 -> []string{"a", "slice"} }`

    lastNode := l.PushBack([]string{"a", "slice"})

Remove the first node `{ 1 -> 2 -> []string{"a", "slice"} }` 
    
    l.Remove(frontNode) 

Remove a middle node `{ 1 -> []string{"a", "slice"} }`
    
    l.Remove(middleNode)

Remove the last node `{ 1 }`
    
    l.Remove(lastNode) 