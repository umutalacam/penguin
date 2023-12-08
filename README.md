# Penguin Credential Store Manager 
Simple credential store manager with high security.

## How to use:

**Get Credential**
```
$ penguin get /ally-bros/google/        
Enter the passphrase for (.credentials/vault):
{
  "type": "directory",
  "name": "google",
  "items": [
    {
      "type": "entity",
      "name": "email",
      "value": "info@allybros.com"
    }
  ]
}

```

**Put Credential**
```
$ penguin put /ally-bros/dev/development-db/credential2 value
 
```
