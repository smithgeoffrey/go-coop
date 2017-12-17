# TODOs

### Error handling

So far I'm ignoring errors largely. Try modifying everything to return, also, error:

    func (c InmemConfig) Get(key string) (string, error) {
        val, ok := c.M[key]
        if ok {
            return val, nil
        } else {
            return "", errors.New("Tried to get a key which doesn't exist")
        }
    }

### Do the testing package


