package rueidis

import "github.com/rueian/rueidis/internal/cmds"

// Commands is an exported alias to []cmds.Completed.
// This allows users to store commands for later usage, for example:
//   client.Dedicated(func(client rueidis.DedicatedClient) error {
//	 	cmds := make(rueidis.Commands, 10)
//	 	for i := 0; i < 10; i++ {
//	 		cmds = append(cmds, client.B().Set().Key(strconv.Itoa(i)).Value(strconv.Itoa(i)).Build())
//	 	}
//	 	for _, resp := range client.DoMulti(ctx, cmds...) {
//	 		if err := resp.Error(); err != nil {
//	 			return err
//	 		}
//	 	}
//	 	return nil
//   })
// However, please know that once commands are processed by the Do() or DoMulti(), they are recycled and should not be reused.
type Commands []cmds.Completed
