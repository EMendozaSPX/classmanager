# Algorithms

## Authentication

### Salt and Hash
```
BEGIN SaltAndHash(password)
    let hash = bcrypt.GenerateFromPassword(password, bcrypt.MinCost) as string
    return hash
END SaltAndHash
```

### Compare Hashes
```
BEGIN CompareHashAndPassword(hash, password)
    let check = bcrypt.CompareHashAndPassword(hash, password) as bool
    IF !check
        return true
    END IF
    return false
END
```

### Hash
