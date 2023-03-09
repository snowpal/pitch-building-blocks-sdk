package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main(jwtToken string, pod request.PodWithParam) error {

	url := "block-pods/%s/copy?keyId=%s&blockId=%s&allTasks=%s&allChecklists=%s&targetKeyId=%s&targetBlockId=%s"
	method := "POST"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang.RouteBlockPodsCopyBlockPod, pod.ID, pod.KeyId, *pod.BlockId, strconv.FormatBool(*pod.AllTasks), strconv.FormatBool(*pod.AllChecklists)), payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
}
